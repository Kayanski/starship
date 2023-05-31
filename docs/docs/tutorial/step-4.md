# **Step 4:** All in one place

## TLDR
We have created some handy example directory for the tutorial you just went though.
You can find it at [examples/tutorial](https://github.com/cosmology-tech/starship/tree/main/examples/getting-started).

You can download the example directory with this [URL](https://download-directory.github.io/?url=https%3A%2F%2Fgithub.com%2Fcosmology-tech%2Fstarship%2Ftree%2Fmain%2Fexamples%2Fgetting-started)

Unzip it into a directory and run the following commands to get started.
You should see
```bash
Makefile
configs/
  starship.yaml
  tiny-starship.yaml
scripts/
  dev-setup.sh
  port-forward.sh
README.md
```

Now inorder to spin a cluster run
```bash
cd getting-started/

# Install dependencies, install startship helm chart, create kind cluster
make setup

# Install the starship instance and run port-forward
make start
# OR, if you are low on resources on local machine
make start-tiny

# Stop the cluster with
make stop
```

Checkout the README.md for more details on the commands, you can even run the commands individually, if some commands are having trouble.

Once you are done, you can all the resources with
```bash
make clean
```