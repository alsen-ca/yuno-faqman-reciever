// Recieves language_translation; meaning either the Tag in english, german or spanish.
// As all of those language_translation[s] must be unique, it searches for an existing Tag with any of those languages.
// Returns the whole tag
func getTag(w http.ResponseWriter, r *http.Request, lang_translation string) {
	let qa Qa = nil
	// Takes lang_translation and searches for a tag equaling it on 3 cases
	tag = perform_cases()
	
	if !tag {
		return 404
	} else {
		return 200, tag
	}
}

func perform_cases() {
	// case 1
	// SELECT * from tag where en_og like `lang_translation`
	// tag = Tag.where(en_og: lang_translation)

	// case 2
	// SELECT * from tag where de_trans like `lang_translation`
	// tag = Tag.where(de_trans: lang_translation)

	// case 3
	// SELECT * from tag where es_trans like `lang_translation`
	// tag = Tag.where(es_trans: lang_translation)

	return tag
}