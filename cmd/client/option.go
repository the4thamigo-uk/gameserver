package main

type option struct {
	link *link
	rel  string
}

type options []*option

func optionsFromLinks(ls links) options {
	opts := options{}
	for _, rel := range relOrder {
		l, ok := ls[rel]
		if ok {
			opts = append(opts,
				&option{
					link: l,
					rel:  rel,
				},
			)
		}
	}
	return opts
}
