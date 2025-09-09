package fail

fail

/*
This is a non-compiling file that has been added to explicitly ensure that CI fails.
It also contains the command that caused the failure and its output.
Remove this file if debugging locally.

go mod operation failed. This may mean that there are legitimate dependency issues with the "go.mod" definition in the repository and the updates performed by the gomod check. This branch can be cloned locally to debug the issue.

Command that caused error:
./godelw check compiles

Output:
Running compiles...
../go/go-dists/go1.25.1/src/internal/runtime/cgroup/cgroup_linux.go:567:13: cannot range over 4 (untyped int constant)
../go/go-dists/go1.25.1/src/internal/runtime/cgroup/cgroup_linux.go:693:18: cannot range over 3 (untyped int constant)
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/internal/runtime/maps/map.go:364:18: cannot range over m.dirLen (variable of type int)
../go/go-dists/go1.25.1/src/internal/runtime/maps/map.go:745:18: cannot range over m.dirLen (variable of type int)
../go/go-dists/go1.25.1/src/internal/runtime/maps/runtime_faststr_swiss.go:33:18: cannot range over abi.SwissMapGroupSlots (untyped int constant 8)
../go/go-dists/go1.25.1/src/internal/runtime/maps/runtime_faststr_swiss.go:64:12: cannot range over abi.SwissMapGroupSlots (untyped int constant 8)
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/runtime/mcleanup.go:554:12: cannot range over wake (variable of type uint32)
../go/go-dists/go1.25.1/src/runtime/mcleanup.go:588:12: cannot range over need (variable of type uint32)
../go/go-dists/go1.25.1/src/runtime/mheap.go:1978:12: cannot range over n - 1 (value of type int)
../go/go-dists/go1.25.1/src/runtime/time.go:236:17: cannot range over 3 (untyped int constant)
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/bytes/bytes.go:1196:13: cannot range over n (variable of type int)
../go/go-dists/go1.25.1/src/bytes/bytes.go:1204:13: cannot range over n - 1 (value of type int)
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/strings/strings.go:1162:13: cannot range over n (variable of type int)
../go/go-dists/go1.25.1/src/strings/strings.go:1170:13: cannot range over n - 1 (value of type int)
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/reflect/iter.go:76:19: cannot range over v.Len() (value of type int)
../go/go-dists/go1.25.1/src/reflect/iter.go:84:19: cannot range over v.Len() (value of type int)
../go/go-dists/go1.25.1/src/reflect/iter.go:140:19: cannot range over v.Len() (value of type int)
../go/go-dists/go1.25.1/src/reflect/iter.go:148:19: cannot range over v.Len() (value of type int)
../go/go-dists/go1.25.1/src/reflect/type.go:2173:19: cannot range over num (variable of type int)
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/slices/iter.go:51:17: cannot range over seq (variable of type iter.Seq[E])
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/os/root_openat.go:125:20: cannot range over 2 (untyped int constant)
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/crypto/internal/fips140/bigmod/nat.go:1204:17: cannot range over size (variable of type int)
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/crypto/internal/fips140/rsa/keygen.go:386:12: cannot range over mr.a - 1 (value of type uint)
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/maps/iter.go:51:20: cannot range over seq (variable of type iter.Seq2[K, V])
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/unique/clone.go:76:12: cannot range over atyp.Len (variable of type uintptr)
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/crypto/x509/verify.go:1490:20: cannot range over pg.parents() (value of type iter.Seq[*policyGraphNode])
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/mime/mediatype.go:258:17: cannot range over len(v) (value of type int)
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/net/http/fs.go:865:19: cannot range over strings.FieldsFuncSeq(v, isSlashRune) (value of type iter.Seq[string])
../go/go-dists/go1.25.1/src/net/http/fs.go:1025:18: cannot range over strings.SplitSeq(s[len(b):], ",") (value of type iter.Seq[string])
../go/go-dists/go1.25.1/src/net/http/server.go:1585:17: cannot range over strings.SplitSeq(v, ",") (value of type iter.Seq[string])
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/regexp/syntax/parse.go:1678:17: cannot range over len(name) (value of type int)
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/net/http/httptest/recorder.go:210:19: cannot range over strings.SplitSeq(k, ",") (value of type iter.Seq[string])
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/testing/testing.go:2706:19: cannot range over strings.SplitSeq(*cpuListStr, ",") (value of type iter.Seq[string])
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
../go/go-dists/go1.25.1/src/go/token/serialize.go:54:17: cannot range over s.tree.all() (value of type iter.Seq[*File])
../go/go-dists/go1.25.1/src/go/token/serialize.go:60:11: cannot infer S (/go/go-dists/go1.25.1/src/slices/slices.go:353:12)
../go/go-dists/go1.25.1/src/go/token/serialize.go:61:11: cannot infer S (/go/go-dists/go1.25.1/src/slices/slices.go:353:12)
-: This application uses version go1.21 of the source-processing packages but runs version go1.25 of 'go list'. It may fail to process source files that rely on newer language features. If so, rebuild the application using a newer version of Go.
Finished compiles
Check(s) produced output: [compiles]

*/
