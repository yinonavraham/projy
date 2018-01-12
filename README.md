# projy
A utility for common tasks in a development project lifecycle

## Usage

```
$ projy <command> [sub-commands...] [arguments...] [options...]
```

### General Options
```
--version    - Print projy version
--help       - Print projy help, either general help or on a specific command or sub-command
--force      - Ignore preconditions or any validation
--dry-run    - Skip non-reversable actions, print expected outcomes instead
--skip-tests - Skip running tests where relevant
--skip-push  - Skip pushing changes to the remote repository
```

### Commands

#### Project Versioning

##### Release Version

```
$ projy version release <version>
```  

Execute a flow for setting a release version, commit, tag and push.    
**Arguments:**  
1. `<version>` - the release version (semantic version)

**Examples:**

```
$ projy version release 1.2.0
$ projy version release 1.2.0 --dry-run
$ projy version release 1.2.0 --force --skip-tests
```

##### Release Candidate
  
```
$ projy version rc <version> <number>
```  

Execute a flow for setting a release candidate version, commit, tag and push.
The actual version is: `<version>-rc<number>` (e.g. `1.2.0-rc1`).  
**Arguments:**  
1. `<version>` - the release version (semantic version) 
1. `<number>` - the rc number

**Examples:**  
```
$ projy version rc 1.2.0 1
  (sets version: 1.2.0-rc1)
$ projy version rc 1.2.0 1 --force --skip-tests
```

##### Snapshot Version

```
$ projy version snapshot <version>
```  

Execute a flow for setting a snapshot version, commit and push.
The actual version is: `<version>[-<qualifier>]-SNAPSHOT`  
**Arguments:**
1. `<version>` - the semantic version  
  
**Options:**  
* `--qualifier=<qualifier>`  
  An additional qualifier to add to the version

**Examples:**  

```
$ projy version snapshot 1.2.0
  (sets version: 1.2.0-SNAPSHOT)
$ projy version snapshot 1.2.0 --qualifier=FOO
  (sets version: 1.2.0-FOO-SNAPSHOT)
```

##### Current Project Version
  
```
$ projy version current
```  

Print the current version of the project. 
Also prints differences between submodules, if any.

### Configuration

#### Project Configuration
To configure `projy` for a specific project, 
create a `.projy.yaml` file in the root of the project.  
File example:
  
```yaml
type: maven
build:
  command: ["mvn", "clean", "install"]
  arguments:
    skip-tests: ["-DskipTests"]
```  


## Implementation

### Commands

#### Project Versioning

##### Release Version

1. Check preconditions: (respects `--force`)
   1. Version format: `<major>.<minor>.<patch>`
   1. Current git branch matches `"release/<version>"` or `"hotfix/<version>"`
1. Set project version: `<version>`
1. Build with unit tests (respects `--skip-tests`, `--dry-run`)
1. Git commit with message: `"Set release version: <version>"` (respects `--dry-run`)
1. Git tag: `<version>` (respects `--dry-run`)
1. Git push, including tag (respects `--dry-run`, `--skip-push`)

##### Release Candidate

1. Check preconditions: (respects `--force`)
   1. Version format: `<major>.<minor>.<patch>`
   1. RC number is a positive integer
   1. Current git branch matches `"release/<version>"` or `"hotfix/<version>"`
1. Set project version: `<version>-rc<number>`
1. Build with unit tests (respects `--skip-tests`, `--dry-run`)
1. Git commit with message: `"Set release candidate version: <version>-rc<number>"` (respects `--dry-run`)
1. Git tag: `<version>-rc<number>` (respects `--dry-run`)
1. Git push, including tag (respects `--dry-run`, `--skip-push`)

##### Snapshot Version
1. Check preconditions: (respects `--force`)
   1. Version format: `<major>.<minor>.<patch>`
   1. Qualifier format: `[a-zA-Z0-9_-]+`
1. Set project version: `<version>[-<qualifier>]-SNAPSHOT`
1. Build with unit tests (respects `--skip-tests`, `--dry-run`)
1. Git commit with message: `"Set snapshot version: <version>"` (respects `--dry-run`)
1. Git push (respects `--dry-run`, `--skip-push`)
