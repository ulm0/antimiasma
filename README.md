
# Antimiasma

This is a Discovery and Mitigation Tool for the Miasma worm.

Additionally, it is an Anti-Worm that combats the Miasma worm by using
the same spread technique in order to spread the vaccination/fix to other
machines and repositories.


### Download

If you don't have `go` installed, ready-to-use downloads are available in the
[Releases section](https://github.com/Team-Rockstars-Security/antimiasma/releases/tag/latest)
on GitHub.


### RECOMMENDED USAGE: Worm Discovery and Mitigation

This shows how to:

- discover infected repositories
- mitigate infected repositories by removing the implants/tasks/scripts

```bash
antimiasma discover ~/Software; # shows a list of infected repositories
antimiasma mitigate ~/Software; # removes the miasma worm in all infected repositories
```

### THE LAST OPTION: Anti-Worm Vaccination

This shows how to:

- discover infected repositories
- mitigate infected repositories by removing the miasma implant
- vaccinate other developer machines by implanting the antimiasma worm

What the antimiasma implant will do:

- if other repositories were infected, it writes the `README-IMPORTANT-ANTIMIASMA.md` audit report into the infection-spreading repository's root folder
- if no other repositories were infected, it will delete itself after a 7 days-since-infection timeout

Please check the [actions/Vaccinate](/actions/Vaccinate.go) and [actions/Immunize](/actions/Immunize.go) files to verify this yourself.
The Vaccinate action has to be executed **once**. The Immunize action is automatically executed on infected machines using the same
spreading technique that the miasma worm uses.

**ONLY** execute this if you have the rights to do this in your organization and are aware of your consequences. **ALWAYS**
coordinate any usage of this tool with your IT Security and NOC/SOC department.

**IMPORTANT NOTE**: SOC teams can monitor ongoing vaccinations of the antimiasma worm within their EDR dashboards. The replaced
[setup.js](/utils/antimiasma/antimiasma.js) file contains the [EICAR test string](https://en.wikipedia.org/wiki/EICAR_test_file) which
triggers every anti-virus program.

```bash
antimiasma vaccinate ~/Software; # vaccinates all infected repositories with antimiasma worm
```


### Testing

A minimal Miasma implant sample is available at [miasma-sample](https://github.com/chris-teamrockstars/miasma-sample).
That repository contains the spreading mechanisms and infiltrated files, but not the actual malware implant.

If you're unsure whether `antimiasma` works on your system, you can clone the `miasma-sample` repository safely
to test the discover and mitigate actions against.

```bash
cd /tmp;

git clone https://github.com/chris-teamrockstars/miasma-sample.git /tmp/miasma-sample;

# should show the /tmp/miasma-sample repository as infected
antimiasma discover /tmp;

# should mitigate the infected repository
antimiasma mitigate /tmp;
```


### Implementation Status

See [utils/miasma](/utils/miasma) and [utils/antimiasma](/utils/antimiasma) folders for implementation details.

| Infected File             | Description                           | Discover | Mitigate | Vaccinate |
|:--------------------------|:--------------------------------------|:--------:|:--------:|:---------:|
| `.claude/settings.json`   | triggers on Claude Code session start | [x]      | [x]      |           |
| `.cursor/rules/setup.mdc` | triggers when opening repo in Cursor  | [x]      | [x]      |           |
| `.gemini/settings.json`   | triggers on Gemini session start      | [x]      | [x]      |           |
| `.vscode/tasks.json`      | triggers when opening repo in VS Code | [x]      | [x]      |           |
| `package.json`            | hijacks `npm test` script             | [x]      | [x]      |           |
| `.github/setup.js`        | the Miasma worm implant               | [x]      | [x]      | [x] [1]   |

The vaccine is implemented in [utils/antimiasma/antimiasma.js](/utils/antimiasma/antimiasma.js)
and relies on `node` as a runtime, same as the original worm.

The `vaccinate` action will discover infected repositories, replace the `setup.js` with the antimiasma
vaccine, and execute `git push` to the currently configured remotes of those repositories. This is the
identical behavior of the original worm. It will auto-mitigate and self-remove itself if after 7 days
no newly infected repositories were discovered on the system.


### Notes

This Anti-Worm has been implemented with an abliterated `qwen3-coder:30b` `@Q8` running
inside the [exocomp](https://github.com/cookiengineer/exocomp) agentic environment.


### Acknowledgements

- RedHat, for disclosing the breach very quickly and correctly
- Mr. Cuddlebug, my cat, for emotional support during implementation
- The Prodigy, for being the best music band of our time
- Person of Interest and Michael Emerson, for the remnant inspiration for this with ICE9
- TeamPCP, the assumed authors of Miasma, for making my weekend more fun


### Reading Material

- https://access.redhat.com/security/vulnerabilities/RHSB-2026-006
- https://safedep.io/miasma-worm-ai-coding-agent-config-injection/


### License

WTFPL

