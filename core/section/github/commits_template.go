// Copyright 2016 Documize Inc. <legal@documize.com>. All rights reserved.
//
// This software (Documize Community Edition) is licensed under
// GNU AGPL v3 http://www.gnu.org/licenses/agpl-3.0.en.html
//
// You can operate outside the AGPL restrictions by purchasing
// Documize Enterprise Edition and obtaining a commercial license
// by contacting <sales@documize.com>.
//
// https://documize.com

package github

const commitsTemplate = `
<div class="section-github-render">
	<!--
	{{if .HasAuthorStats}}
		<div class="heading">Contributors</div>
		<p>
			There
			{{if eq 1 .NumContributors}}is{{else}}are{{end}}
			{{.NumContributors}}
			{{if eq 1 .NumContributors}}contributor{{else}}contributors{{end}}
			across {{.RepoCount}}
			{{if eq 1 .RepoCount}} repository. {{else}} repositories. {{end}}
		</p>

		<table class="github-table">

			<thead>
				<tr>
				<th class="title">Contributors</th>
				<th></th>
				</tr>
			</thead>

			<tbody>
				{{range $stats := .AuthorStats}}
					<tr>
						<td class="no-width">
							<img class="github-avatar" alt="@{{$stats.Author}}" src="{{$stats.Avatar}}" />
						</td>
						<td>
							<div class="contributor-name">{{$stats.Author}}</div>
							<div class="contributor-meta">
								{{if gt $stats.OpenIssues 0}}
									assigned {{$stats.OpenIssues}}
									{{if eq 1 $stats.OpenIssues}} issue {{else}} issues {{end}}
								{{end}}
								{{if gt $stats.ClosedIssues 0}}
									 &middot; {{$stats.ClosedIssues}} closed
								{{end}}
								{{if gt $stats.CommitCount 0}}
									{{if gt $stats.OpenIssues 0}} &middot; {{end}}
									{{if gt $stats.ClosedIssues 0}} &middot; {{end}}
									made {{$stats.CommitCount}}
									{{if eq 1 $stats.CommitCount}} commit {{else}} commits {{end}}
									on {{len $stats.Repos}} {{if eq 1 (len $stats.Repos)}} branch {{else}} branches {{end}}
									{{range $repo := $stats.Repos}}	&middot; {{$repo}} {{end}}
								{{end}}
							</div>
						</td>
					</tr>
				{{end}}
			</tbody>
		</table>
	{{end}}
	-->

	{{if .HasCommits}}
		<table class="github-table" style="width: 100%;">
			<thead>
				<tr>
				<th class="title">Commits <span>&middot; {{len .BranchCommits}} commits by {{.NumContributors}} contributors</span>
				</th>
				<th></th>
				</tr>
			</thead>
			<tbody>
				{{range $commit := .BranchCommits}}
					<tr>
						<td>
							<a href="{{$commit.URL}}">{{$commit.Message}}</a>
							<span class="data"> {{$commit.Branch}}</span>
						</td>
						<td class="right-column">
							<div class="contributor-meta">
								<img class="github-avatar" alt="@{{$commit.Name}}" src="{{$commit.Avatar}}" />
								{{$commit.Name}}
								&middot; {{$commit.Date}}
							</div>
						</td>
					</tr>
				{{end}}
			</tbody>
		</table>
	{{end}}

</div>
`
