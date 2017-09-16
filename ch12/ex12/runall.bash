#!/bin/bash/
cd `dirname $0`

cd fetch
go build -o fetch

cd ../search
go build -o search

cd ../

# ./search/search &
# ./fetch/fetch 'http://localhost:12345/search'
# ./fetch/fetch 'http://localhost:12345/search?em=test1@test.ne.jp&em=test2@test.ne.jp'
# ./fetch/fetch 'http://localhost:12345/search?em=@test.ne.jp'
# ./fetch/fetch 'http://localhost:12345/search?cn=12345678901234'
# ./fetch/fetch 'http://localhost:12345/search?cn=12345'
# ./fetch/fetch 'http://localhost:12345/search?pn=1234567'
# ./fetch/fetch 'http://localhost:12345/search?pn=12345'