---
title: "git-qtk"
template: "article"
path: "/git-qtk/"
meta:
  description: "git-qtk (Git Query Toolkit) is a command-line tool that lets you query Git repository metadata—like commits, authors, and files—using a simple, declarative SQL-like language in YAML files."
  image: ""
---

# git-qtk: Query Your Git Repositories Declaratively

git-qtk (Git Query Toolkit) is a command-line tool I built as my thesis project. It makes it easy to explore and analyze Git repository history by writing readable YAML queries that feel like SQL—but tailored specifically for Git data.

Instead of chaining complex `git log` commands or writing custom scripts, you define what you want in a clear, version-controllable file: select fields, filter commits, join authors with their commits, group results, and even use JavaScript for custom logic. The tool handles the rest efficiently, producing clean tabular output.

Below are real-world examples that show what git-qtk can do.

## Listing Recent Commits

A simple query to get the latest commits with author names:

```yaml
from: commit c; author a
select: short(c.sha); a.name; c.message
where: c.author == a.id
order: c.date desc
limit: 10
```

**Example output**:

```
┌────────────────┬──────────────────┬────────────────────────────────────────┐
│ short(c.sha)   │ a.name           │ c.message                              │
├────────────────┼──────────────────┼────────────────────────────────────────┤
│ a1b2c3d        │ John Doe         │ Fix bug in parser                      │
│ e4f5g6h        │ Jane Smith       │ Add new query examples                 │
│ ...            │ ...              │ ...                                    │
└────────────────┴──────────────────┴────────────────────────────────────────┘
```

## Top Contributors by Commit Count

Who’s been the most active?

```yaml
from: commit c; author a
select: a.name; count(c.sha) as commits
where: c.author == a.id
group: a.name
order: commits desc
limit: 5
```

**Example output**:

```
┌──────────────────┬─────────┐
│ a.name           │ commits │
├──────────────────┼─────────┤
│ Alice Johnson    │ 142     │
│ Bob Lee          │ 98      │
│ John Doe         │ 76      │
│ Jane Smith       │ 54      │
│ Eve Brown        │ 31      │
└──────────────────┴─────────┘
```

## Most Frequently Changed Files

Which files get touched the most?

```yaml
from: file f; commit c
select: f.path; count(c.sha) as changes
where: c.changes.indexOf(f.path) >= 0
group: f.path
order: changes desc
limit: 10
```

**Example output** (from the git-qtk repo itself):

```
┌────────────────────┬─────────┐
│ f.path             │ changes │
├────────────────────┼─────────┤
│ lib/parser.js      │ 45      │
│ bin/main.js        │ 23      │
│ README.md          │ 18      │
│ lib/runner.js      │ 15      │
│ ...                │ ...     │
└────────────────────┴─────────┘
```

## Filtering Specific File Types in a Directory

How many times were JavaScript files in `bin/` changed?

```yaml
from: file f; commit c
select: f.path; count(c.sha) as changes
where: c.changes.indexOf(f.path) >= 0 && 
       f.path.endsWith('.js') && 
       f.path.startsWith('bin/')
group: f.path
order: changes desc
```

**Real output** when run on the git-qtk repository:

```
┌────────────────┬──────────────┐
│ f.path         │ changes      │
├────────────────┼──────────────┤
│ bin/runtime.js │ 26           │
│ bin/main.js    │ 23           │
│ bin/ctt.js     │ 12           │
└────────────────┴──────────────┘
```

## JavaScript Expressions for Custom Logic

Need something more specific? Use inline JavaScript.

Commits in the last 30 days with formatted dates:

```yaml
from: commit c
select: short(c.sha); c.message; new Date(c.date).toLocaleDateString()
where: c.date > now() - (30*24*60*60*1000) /* 30 days */
order: c.date desc
```

Or fancy author titles:

```yaml
select: "Dr. " + a.name; count(c.sha) as commits
```

## Installation and Quick Start

Install globally via npm (requires Node.js and Git):

```bash
npm install https://github.com/imdonix/git-qtk --global
```

Test it:

```bash
git-qtk -v
```

Run any query:

```bash
git-qtk -s your-query.yaml -r /path/to/repo
# or a remote repo
git-qtk -s your-query.yaml -r https://github.com/torvalds/linux
```

The tool works with local paths or direct Git URLs—it clones temporarily if needed.

## Thanks

Thank you for reading about git-qtk! This was my thesis project, exploring how declarative querying can make Git history analysis more accessible and powerful.

The full source code, query syntax details, and more are on GitHub:
<a href="https://github.com/imdonix/git-qtk">github.com/imdonix/git-qtk</a>

If you try it out, find it useful, or have ideas for new queries/features, open an issue or reach me on X:
<a href="https://x.com/imdonix">@imdonix</a>

Happy querying!
