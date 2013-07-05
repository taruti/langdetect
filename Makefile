all: \
	autogenerated_iso639.go \
	autogenerated_iso639_en_names.go \
	autogenerated_iso639_nat_names.go \
	autogenerated_cp1252.go \
	test_data/wikipedia/auto_en.txt \
	autogenerated_trigrams_en.go \


P=langdetect

autogenerated_trigrams_en.go:
	echo This needs hand tweaking
	false

autogenerated_common.words.go: common.words.*
	echo package $P > $@
	echo >>$@
	echo "var commonWords = map[string]Language{" >>$@
	for f in common.words.*; do L=$$(echo $$f | cut -d. -f3 | sed 's/./\U&/'); for w in `egrep -v '^#' $$f`; do echo \"$$w\":$$L, >> $@; done; done
	echo "}">>$@

autogenerated_iso639.go:
	echo package $P > $@
	echo >>$@
	echo 'var (' >>$@
	cat iso639.raw | cut -f1 | perl -pe "\$$c=\$$c+0;s/^(\w)(\w)/ \u\\1\\2 = Language{[2]byte{'\\1','\\2'},\$$c,0}/;\$$c++" >>$@
	echo ')' >>$@

autogenerated_iso639_en_names.go:
	echo package $P > $@
	echo >>$@
	echo "var langEnglish = [...]string{" >>$@
	cat iso639.raw | cut -f1,3 | perl -pe '$$c=$$c+0;s/^(\w\w)\t(.+)/\u$$c:"$$2",/;$$c++' >>$@
	echo "}">>$@

autogenerated_iso639_nat_names.go:
	echo package $P > $@
	echo >>$@
	echo "var langNative = [...]string{" >>$@
	cat iso639.raw | cut -f1,4 | perl -pe '$$c=$$c+0;s/^(\w\w)\t(.+)/\u$$c:"$$2",/;$$c++' >>$@
	echo "}">>$@

autogenerated_cp1252.go: bestfit1252.txt
	echo package $P > $@
	echo >>$@
	echo "var cp1252 = smallhalfencoding{" >>$@
	egrep '^0x[[:alnum:]]{2}[[:space:]]' bestfit1252.txt|cut -f2|tail -128|sed -e 's/\r//;s/$$/,/' >>$@
	echo "}">>$@

test_data/wikipedia/auto_en.txt: test_data/wikipedia_languages.txt
	mkdir -p test_data/wikipedia
	for L in `cat test_data/wikipedia_languages.txt`; do echo Fetching wikipedia front page in $$L;lynx -dump -nolist https://$${L}.wikipedia.org/ >test_data/wikipedia/auto_$${L}.txt ; done

clean:
	rm -f autogenerated_* test_data/*/auto_*

