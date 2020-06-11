---
id: docker
title: Docker
---

For running in docker just use the following command:
```
docker run --rm --name="commitsar" -w /src -v "$(pwd)":/src aevea/commitsar
```

Make sure to load the working directory where `.git` folder is present. Commitsar will not work without the `.git` folder.