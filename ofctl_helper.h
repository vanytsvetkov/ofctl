#include <stdlib.h>

#include <getopt.h>

#include "command-line.h"

int
ofctl_run(int argc, char *argv[]);

const struct ovs_cmdl_command *get_all_commands(void);
void parse_options(int argc, char *argv[]);
