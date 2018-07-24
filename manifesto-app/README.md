# Manifesto

This folder contains the source code for the Kyma Manifesto page. Manifesto is a statement of ideals and intentions the Kyma community wants to follow during the development of Kyma. It is published [here](https://kyma-project.github.io/community/).

## Docker image build and run

To build and run the Docker image, follow these steps:

```
docker build -t kyma-manifesto .
docker run --rm -p 8080:80 kyma-manifesto
open http://localhost:8080 in a browser
```

## Publish on GitHub Pages

To publish Manifesto from the `master` branch on GitHub Pages, run the following script:

```bash
./publish-gh-page.sh
```
