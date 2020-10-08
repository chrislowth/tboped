# TbOpEd

TurBonomic OPerator EDitor.

*Status*: Experimental, hopefully useful, unsupported.

This "personal project" tool is an attempt to make editting the Turbonomic Operator chart YAML file easier and safer.

It runs as a text-only "curses" style editor which you need to run on a x86_64 linux system.

This version of the tool allows you to:

- Change the version tag.
- Toggle feature "enable" flags from true to false and/or back.
- Add features listed in the "values.yaml" file that are not mentioned the chart file.

Other features may come in due course.

## Environment considerations

The tool accesses the following resources...

1: The helm chart itself (usually: charts_v1alpha1_xl_cr.yaml) - the file being editted.
2: The chart definition file (usually: charts_v1alpha1_xl_crd.yaml) for the purpose of presenting a documentary hint about the role of features as they are selected.
3: The "values.yaml" file help in the file system of the operator POD. This is collected using "kubectl exec .. cat ..".

For these reasons, the tool must be run from a location that can access all three sources. For OVA installed instance, the perfect place is on the instance itself.

## Installation

1. Download the tar file containing the compiled binary from the GITHUB repository at ... https://turbonomic-emea-cs-bucket.s3.eu-west-2.amazonaws.com/tboped/tboped.tgz
2. Copy the tar file to the XL instance's `/tmp` directory.
3. Unzip it into a directory that commands can be run from. Eg: $HOME/bin

## Execution

Just run "tboped" and it will pick up the files from the default locations..

- /opt/turbonomic/kubernetes/operator/deploy/crds/charts_v1alpha1_xl_cr.yaml
- /opt/turbonomic/kubernetes/operator/deploy/crds/charts_v1alpha1_xl_crd.yaml
- "kubectl" on the linux path.

If you wish to edit a different `_cr.yaml` file or use descriptions from a different `_crd.yaml` file then you can specific alternative paths on the command line. The first argument is the name of the `_cr.yaml` file to be editted and the second (if specified) is the alternative `_crd.yaml`.

In most cases you can just run it without arguments.

After making the edits, review the file and (if you're happy with it). run "kubectl apply -f {filename}" to apply the changes.
