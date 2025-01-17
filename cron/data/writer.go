// Copyright 2021 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"

	"github.com/jszwec/csvutil"

	"github.com/ossf/scorecard/repos"
)

type repoEntry struct {
	Repo     string `csv:"repo"`
	Metadata string `csv:"metadata"`
}

func SortAndAppend(out io.Writer, newRepos []repos.RepoURL) error {
	iter, err := MakeIterator()
	if err != nil {
		return fmt.Errorf("error during MakeIterator: %w", err)
	}

	oldRepos := make([]repoEntry, 0)
	for iter.HasNext() {
		repo, err := iter.Next()
		if err != nil {
			return fmt.Errorf("error during iter.Next: %w", err)
		}
		repoentry := repoEntry{
			Repo:     repo.URL(),
			Metadata: repo.Metadata,
		}
		oldRepos = append(oldRepos, repoentry)
	}
	for _, newRepo := range newRepos {
		repoentry := repoEntry{
			Repo:     newRepo.URL(),
			Metadata: newRepo.Metadata,
		}
		oldRepos = append(oldRepos, repoentry)
	}
	sort.SliceStable(oldRepos, func(i, j int) bool {
		return oldRepos[i].Repo < oldRepos[j].Repo
	})

	csvWriter := csv.NewWriter(out)
	enc := csvutil.NewEncoder(csvWriter)
	if err := enc.Encode(oldRepos); err != nil {
		return fmt.Errorf("error during Encode: %w", err)
	}
	csvWriter.Flush()
	return nil
}
