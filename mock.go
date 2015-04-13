package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/shurcooL/Go-Package-Store/presenter"
)

func mockHandler(w http.ResponseWriter, req *http.Request) {
	if err := loadTemplates(); err != nil {
		fmt.Fprintln(w, "loadTemplates:", err)
		return
	}

	//started := time.Now()

	CommonHat(w)
	defer CommonTail(w)

	io.WriteString(w, `<div id="checking_updates"><h2 style="text-align: center;">Checking for updates...</h2></div>`)
	io.WriteString(w, `<div id="no_updates" style="display: none;"><h2 style="text-align: center;">No Updates Available</h2></div>`)
	defer io.WriteString(w, `<script>document.getElementById("checking_updates").style.display = "none";</script>`)

	flusher := w.(http.Flusher)
	flusher.Flush()

	notifier := w.(http.CloseNotifier)
	go func() {
		<-notifier.CloseNotify()

		//fmt.Println("Exiting, since the HTTP request was cancelled/interrupted.")
		//close(updateRequestChan)
	}()

	//fmt.Printf("Part 1: %v ms.\n", time.Since(started).Seconds()*1000)

	//fmt.Printf("Part 2: %v ms.\n", time.Since(started).Seconds()*1000)

	updatesAvailable := 0

	for _, repoPresenter := range repoPresenters {
		//started2 := time.Now()

		time.Sleep(time.Second)

		updatesAvailable++
		writeRepoHtml2(w, repoPresenter)

		flusher.Flush()

		//fmt.Printf("Part 2b: %v ms.\n", time.Since(started2).Seconds()*1000)
	}

	if updatesAvailable == 0 {
		io.WriteString(w, `<script>document.getElementById("no_updates").style.display = "";</script>`)
	}

	//fmt.Printf("Part 3: %v ms.\n", time.Since(started).Seconds()*1000)
}

func writeRepoHtml2(w http.ResponseWriter, repoPresenter interface{}) {
	err := t.Execute(w, repoPresenter)
	if err != nil {
		log.Println("t.Execute:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var repoPresenters = []map[string]interface{}{
	{
		"Repo": map[string]interface{}{
			"ImportPathPattern": (string)("github.com/gopherjs/gopherjs/..."),
			"ImportPaths":       (string)("github.com/gopherjs/gopherjs/compiler\ngithub.com/gopherjs/gopherjs/compiler/astutil\ngithub.com/gopherjs/gopherjs/nosync\ngithub.com/gopherjs/gopherjs/tests\ngithub.com/gopherjs/gopherjs/compiler/analysis\ngithub.com/gopherjs/gopherjs/compiler/typesutil\ngithub.com/gopherjs/gopherjs/js\ngithub.com/gopherjs/gopherjs/build\ngithub.com/gopherjs/gopherjs/compiler/filter\ngithub.com/gopherjs/gopherjs/gcexporter\ngithub.com/gopherjs/gopherjs\ngithub.com/gopherjs/gopherjs/compiler/prelude"),
			"GoPackages": [12]interface{}{
				map[string]interface{}{
					"Bpkg": map[string]interface{}{"ImportPath": (string)("github.com/gopherjs/gopherjs/compiler")},
					"Dir": map[string]interface{}{
						"Repo": map[string]interface{}{
							"VcsLocal":  map[string]interface{}{"LocalRev": (string)("aff1494482d249eb0a3803abbd434f4c9143a3de")},
							"VcsRemote": map[string]interface{}{"RemoteRev": (string)("87bf7e405aa3df6df0dcbb9385713f997408d7b9")},
						},
					},
				},
			},
		},
		"HomePage": (*template.URL)(newTemplateUrl("https://github.com/gopherjs/gopherjs")),
		"Image":    (template.URL)("https://avatars.githubusercontent.com/u/6654647?v=3"),
		"Changes": ([]presenter.Change)([]presenter.Change{
			(presenter.Change)(presenter.Change{
				Message: (string)("improved reflect support for blocking functions"),
				Url:     (template.URL)("https://github.com/gopherjs/gopherjs/commit/87bf7e405aa3df6df0dcbb9385713f997408d7b9"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("small cleanup"),
				Url:     (template.URL)("https://github.com/gopherjs/gopherjs/commit/77a838f965881a888416bae38f790f76bb1f64bd"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("replaced js.This and js.Arguments by js.MakeFunc"),
				Url:     (template.URL)("https://github.com/gopherjs/gopherjs/commit/29dd054a0753760fe6e826ded0982a1bf69f702a"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
		}),
	},

	{
		"Repo": map[string]interface{}{
			"ImportPathPattern": (string)("golang.org/x/image/..."),
			"ImportPaths":       (string)("golang.org/x/image/bmp\ngolang.org/x/image/cmd/webp-manual-test\ngolang.org/x/image/draw\ngolang.org/x/image/riff\ngolang.org/x/image/tiff\ngolang.org/x/image/tiff/lzw\ngolang.org/x/image/vp8\ngolang.org/x/image/vp8l\ngolang.org/x/image/webp\ngolang.org/x/image/webp/nycbcra"),
			"GoPackages": [10]interface{}{
				map[string]interface{}{
					"Bpkg": map[string]interface{}{"ImportPath": (string)("github.com/gopherjs/gopherjs/compiler")},
					"Dir": map[string]interface{}{
						"Repo": map[string]interface{}{
							"VcsLocal":  map[string]interface{}{"LocalRev": (string)("b57ddf1b6833a418ae37c0b2e3570507379abfa3")},
							"VcsRemote": map[string]interface{}{"RemoteRev": (string)("f510ad81a1256ee96a2870647b74fa144a30c249")},
						},
					},
				},
			},
		},
		"HomePage": (*template.URL)(newTemplateUrl("http://golang.org/x/image/bmp")),
		"Image":    (template.URL)("https://avatars.githubusercontent.com/u/4314092?v=3"),
		"Changes": ([]presenter.Change)([]presenter.Change{
			(presenter.Change)(presenter.Change{
				Message: (string)("draw: generate code paths for image.Gray sources."),
				Url:     (template.URL)("https://github.com/golang/image/commit/f510ad81a1256ee96a2870647b74fa144a30c249"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
		}),
	},

	{
		"Repo": map[string]interface{}{
			"ImportPathPattern": (string)("github.com/influxdb/influxdb/..."),
			"ImportPaths":       (string)("github.com/influxdb/influxdb\ngithub.com/influxdb/influxdb/admin\ngithub.com/influxdb/influxdb/client\ngithub.com/influxdb/influxdb/cmd/influxd\ngithub.com/influxdb/influxdb/graphite\ngithub.com/influxdb/influxdb/influxql\ngithub.com/influxdb/influxdb/cmd/influx\ngithub.com/influxdb/influxdb/httpd\ngithub.com/influxdb/influxdb/statik\ngithub.com/influxdb/influxdb/messaging\ngithub.com/influxdb/influxdb/raft\ngithub.com/influxdb/influxdb/collectd"),
			"GoPackages": [12]interface{}{
				map[string]interface{}{
					"Bpkg": map[string]interface{}{"ImportPath": (string)("github.com/influxdb/influxdb")},
					"Dir": map[string]interface{}{
						"Repo": map[string]interface{}{
							"VcsLocal":  map[string]interface{}{"LocalRev": (string)("325f613a5dca621460cc06d47ac7bf327a9b0e73")},
							"VcsRemote": map[string]interface{}{"RemoteRev": (string)("6f398c1daf88fe34faede69f4404a334202acae8")},
						},
					},
				},
			},
		},
		"HomePage": (*template.URL)(newTemplateUrl("https://github.com/influxdb/influxdb")),
		"Image":    (template.URL)("https://avatars.githubusercontent.com/u/5713248?v=3"),
		"Changes": ([]presenter.Change)([]presenter.Change{
			(presenter.Change)(presenter.Change{
				Message: (string)("Add link to \"How to Report Bugs Effectively\""),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/6f398c1daf88fe34faede69f4404a334202acae8"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Update CONTRIBUTING.md"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/37fa6056009dd4e84e9852ec50ce747e22375a99"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Update CONTRIBUTING.md"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/87a6a8f15a13c5bf0ac60608edc1be570e7b023e"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Add note about requiring distro details"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/901f91dc9559bebddf9b49607eac4ffd5caa4158"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Correct typo in change log"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/8eefdba0d3ef3ab5a408073ae275d495b67c9535"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Correct markdown for URL"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/41688ea6af78d45d051c7f6ac24a6468d36b9fad"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Update with PR1744"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/db09b20d199c973a209e181c9e2f890969bd0b57"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Merge pull request #1770 from kylezh/dev"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/a7c0d71d9ccadde17e7aa5cbba538b4a99670633"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Merge pull request #1787 from influxdb/measurement_batch_in_series"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/40479784e2bd690b9021ec730287c426124230dd"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Store Measurement commands in batches"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/a5749bebfb40239b8fd7b25d2ab1aa234c31c6b2"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Merge pull request #1786 from influxdb/remove-syslog"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/2facd6158620e86262407ae3c4c131860f6953c5"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Merge pull request #1785 from influxdb/1784"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/4a5fdcc9ea3bf6dc178f45758332b871e45b93eb"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Fix urlgen to work on Ubuntu"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/666d09367690627f9c3212c1c25c566416c645da"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Remove unused syslog.go"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/06bfd9c496becacff404e6768e7c0fd8ce9603c2"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Fix timezone abbreviation."),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/06eac99c230dcc24bee9c3e1c1ef01725ce017ad"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Merge pull request #1782 from influxdb/more_contains_unit_tests"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/fffbcf3fbe953e03e69ac1d22c142ecd6b3aba3b"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("More shard \"contains\" unit tests"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/ec93341f3fddd294f404fd1469fb651d4ba16e4c"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Update changelog for rc6 release"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/65b4d1a060883a5901bd7c40492a3345d2eabc77"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Merge pull request #1781 from influxdb/single_shard_data"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/5889b12832b2e43424951c92089db03f31df1078"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Refactor shard group time bound checking"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/05d630bfb8041362c89249e3e6fabe6261cecc66"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
			(presenter.Change)(presenter.Change{
				Message: (string)("Fix error when alter retention policy"),
				Url:     (template.URL)("https://github.com/influxdb/influxdb/commit/9f8639ded8778a270cc99cf2d9ee1a09f635d67d"),
				Comments: (presenter.Comments)(presenter.Comments{
					Count: (int)(0),
					Url:   (template.URL)(""),
				}),
			}),
		}),
	},
}

func newTemplateUrl(v template.URL) *template.URL { return &v }
