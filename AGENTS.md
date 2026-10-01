[English](AGENTS.md) · [中文](AGENTS.zh.md)

# be/go-infra

The AI guide to developing this shell. What it hosts and how to deploy it: `BRICKKIT.md`. Members, configuration and port: `component.yaml`.

## Code map

| Path | Owns |
|---|---|
| `main.go` | The whole program: `shell.Main("be-go-infra", shell.Registry{...})`, one entry per member mapping its ID to its `module.New` |
| `go.mod` | The be-sdk-go version and, per member, its Go module at the exact version in `shell.members` |
| `go.sum` | Checksums for `go.mod` |
| `Dockerfile` | Two-stage build; the runtime stage is alpine with `/bin/sh` and `wget`, and holds `component.yaml` next to the binary |
| `component.yaml` | Members (`shell.members`), the shell's own configuration keys, port 8224, health check |
| `BRICKKIT.md` | What a project must know to deploy the shell |

| Task | Start here | Then |
|---|---|---|
| Add a member | `shell.members` in `component.yaml` | its `module.New` in the Registry in `main.go`, its module in `go.mod`, the Shell declaration in `BRICKKIT.md` |
| Move a member to a new version | `shell.members` | `go.mod`, then bump `metadata.version`, rebuild and release (Release below) |
| Change the shell's own configuration | `configSchema` in `component.yaml` | Configuration in `BRICKKIT.md` |

## Build and test

```bash
go build -o /dev/null ./...        # compiles the shell and every registered member
brickkit build be/go-infra           # the image, tagged with metadata.version; run from the project root, with the shell in its brickkit.yaml
make docs-check ID=be/go-infra     # manifest and documents (project root; a project-level lint cannot be scoped to one component)
```

Success: `go build` prints nothing; `brickkit build` records in the image the member versions it compiles in, which `brickkit up` checks; the container turns healthy and `GET /healthz` on port 8224 answers 200.

The shell has no tests of its own: the launcher and its failure contract are tested in be-sdk-go, and each member is tested in its own repository.

## Release

The shell is released from the root of this repository, like a component:

```bash
git push origin main                          # the release check needs the commit on the remote
brickkit release --notes-file <notes file>    # tag <version> (e.g. 1.0.0), annotated and pushed; keep the notes file outside this directory
```

The tag is the bare version. There is no `v` tag: nothing imports the shell as a Go module. Adding a member or moving one to a new version is therefore a commit and a release here first; the assembly project then commits the new submodule pointer and runs `brickkit upgrade be/go-infra@<version>` and `brickkit build be/go-infra`.

## Design decisions

- The shell is a repository of its own, so every assembly project that merges these components builds the same image from the same member list; be-assembly-standard checks it out as a Git submodule at `shell/be/go-infra` (its decision 0108). Which of the compiled-in members a deployment hosts is still chosen in that project's deploy file.
- One shell is one image with one member list: `shell.members` names the exact member versions compiled in, and the Registry in `main.go` lists exactly those members.
- All launcher logic lives in the SDK (`shell.Main`), so this directory stays a list of members and cannot grow logic of its own.
- The shell never runs migrations: brickKit runs each member's migration from the member's own image before the shell starts.

## Pitfalls

| Never | Symptom | Why |
|---|---|---|
| Leave `shell.members` empty and expect the manifest to load | `brickkit lint` and `brickkit add` refuse it: `MANIFEST_INVALID`, "a shell must list at least one component compiled into it"; the shell cannot be added to the project or built with `brickkit build` | brickKit requires at least one compiled-in member. A deployment that hosts none of them is still legal: it is chosen in the deploy file, and the SDK then starts with `BRICKKIT_SERVED_MEMBERS_CONFIG` set to `[]` and serves only `/healthz` |
| Drop `COPY component.yaml` from the `Dockerfile`, or change the working directory | The container exits at once: "读外壳自己的 component.yaml 失败" | `shell.Main` reads its own port (8224) from `deployment.port` in `./component.yaml` |
| Run the shell without `AUTHZ_BUNDLE_URL` or `IAM_JWKS_URL` | The process exits at start with a message naming the missing key ("缺少权限地址配置") | The Go shell loads the permission bundle and the signing keys once for the whole process; without them every member would answer `403` / `503` on every protected route, so it refuses to start |
| Read the 503s of protected routes as a broken shell while infra/authz is not hosted | Members answer `503` on protected routes; `/healthz` stays 200 | `AUTHZ_BUNDLE_URL` points at infra/authz, which this shell hosts itself; until it is hosted and its bundle has loaded once, no permission check can pass |

## Before changing code

1. Adding, removing or moving a member changes `shell.members`, the Registry in `main.go` and `go.mod` together, and the Shell declaration in `BRICKKIT.md`.
2. Every member has a row in the assembly project's `registry/schemas.tsv`, and that project's `make db-init` grants the member's role to `shell_go_infra` and refuses a member without one.
3. A changed member list or member version means a new `metadata.version` for the shell, a rebuild and a release from this repository; `brickkit up` stops a stale image with `IMAGE_STALE`.
4. Code here only registers members: no route, handler, query or call between members.
5. Run `make docs-check ID=be/go-infra` (from the project root; a plain `brickkit lint` there lints the whole project) and the build in Build and test before committing.

<!-- brickkit:managed:begin lang=en -->
<!-- maintained by brickkit (init, add, remove, upgrade, skills update): edits between these markers are overwritten -->

## BrickKit

This is a BrickKit component: `component.yaml` is all the platform reads. The rules it relies on:

- `configSchema` keys are the environment variable names the code reads. Never use a reserved name: `COMPONENT_ID`, `COMPONENT_VERSION`, `PORT`, `BRICKKIT_SERVED_MEMBERS`, `BRICKKIT_SERVED_MEMBERS_CONFIG`, or any `*_ENDPOINT`.
- Dependencies are exact versions. A dependency's address arrives as `<ID>_ENDPOINT`; an optional dependency that is absent has no variable at all, so read it with a fallback.
- `/healthz` checks only this process, never a dependency. The migration command runs from the same image and must fail on an argument it does not know.
- `BRICKKIT.md` travels to every project that uses this component and is read there without the repository: keep it in step with the code, with no relative links.
- Release: raise `metadata.version`, commit, push, `brickkit release`. `brickkit lint` checks the manifest and these docs.
- The full rules are in the `brickkit-component` skill (`.claude/skills/brickkit-component/SKILL.md` at the root of the project or repository where skills are installed; `brickkit skills update` installs it); for flags ask `brickkit <command> --help`.
<!-- brickkit:managed:end -->
