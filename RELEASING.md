# Releasing Moment

## Preparing A Release

1. Update the changelog and the `Version` constant in `hooks.go`.
2. Run `gofmt -l .`, `go vet ./...` and `go test ./...`, including the time-zone
   matrix of `.github/workflows/timezones.yml`.
3. Review and commit the changes.
4. Ensure the protected source branch is green.
5. Create and push the matching version tag, for example `v2.31.0`.

## Automated Validation

The release workflow runs for version tags such as `v2.31.0`. It:

1. Verifies that the tag matches the `Version` constant.
2. Runs the formatting, vet and test checks.
3. Builds `cmd/momentcase`.
4. Updates the `master` branch to point at the release tag.

Release runs are serialized so each tag is compared with the release that
completed immediately before it. For a transient failure, use GitHub Actions to
re-run the failed jobs or the complete workflow; the re-run retains the original
tag context.

## Recovering From A Bad Release Tag

If validation fails for a version tag, correct the release commit and replace
the tag. For example, for `v2.31.0`:

```bash
git push origin --delete v2.31.0
git tag --delete v2.31.0
git tag v2.31.0
git push origin v2.31.0
```

Do not force-update a protected tag. If repository policy prevents deleting or
moving the tag, ask a repository administrator for assistance or release a new
version instead.
