---
weight: 1
bookFlatSection: true
title: "Documenting"
---
# How to Document
## Style Guide
I would like to orient on the [Google Developer Documentation Style Guide](!https://developers.google.com/style)

## Hugo
We throw the documentation in here, which Hugo takes and make a nice page from. There is no big difference in how to write it, just that we directly need to think about where to place what. However, we will have a fast browsable and clean page instead for spread around or to long md's.

## Pages
To create new pages/files go to /docs/Alexandria/content/docs/
You can structure files in folders, if a folder should represent a overview page you need to setup a _index.md file.
Else just write a new md file, don't forget the header:
```md
---
weight: 1 # order of the page
bookFlatSection: true # if a navigation on the right site for the page should be shown
title: "Documenting" # title
---
```

## Images
If you want to use a picture, these can be placed within the content folder or better within the static folder. If placed in static you can link to it from everywhere as follow:
```md
![Hihlevel Arch](/Highlevel.png)
```

## Specify Code Block Types
When explain something with code references or so please ensure to use within markdown the possibility to highlight for specific languages so it turns
```
```go
func main()
```


```go
func main() {
	fmt.Println("Hello you!")
}
```
or
```
```sh
kubectl delete -f https://raw.githubusercontent.com/arangodb/kube-arangodb/${VERSION}/manifests/arango-crd.yaml
```


```sh
kubectl delete -f https://raw.githubusercontent.com/arangodb/kube-arangodb/${VERSION}/manifests/arango-crd.yaml

```

## About this Theme
This theme has some nice add ons which you can discover [here](https://themes.gohugo.io//theme/hugo-book/) and find example for it on the [github](https://github.com/alex-shpak/hugo-book/tree/master/exampleSite/content) repo.