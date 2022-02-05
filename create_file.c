// #include <stdio.h>
// #include <signal.h>
// #include <unistd.h>
// #include <stdlib.h>
// #include <string.h>
// #include <sys/types.h>
// #include <sys/wait.h>
// #include <sys/stat.h>
// #include <errno.h>
// int main ()
// {
//    size_t buflen = 0;
//    char **argv;
//    argv[0]= "ls";
//    argv[1]= "-ltr";
// printf("Enter username: ");

   
//     for(char **argv_scan=argv; *argv_scan != NULL; ++argv_scan) {
//         buflen += strlen(*argv_scan) + 1;
//     }
//     char *cmdline;
//     cmdline = calloc(sizeof(unsigned char), buflen);
//     cmdline += strlen(cmdline);
//     for(; *argv != NULL && (0 != strcmp(*argv, "&")); ++argv) {
//         strcat(cmdline, *argv);
//         cmdline += strlen(cmdline);
//         *cmdline = ' ';
//         ++cmdline;
//     }
//     system(cmdline);
//     free(cmdline);

//    return 0;
// }

/******************************************************************************

                            Online C Compiler.
                Code, Compile, Run and Debug C program online.
Write your code in this editor and press "Run" button to compile and execute it.

*******************************************************************************/

#include <stdlib.h>
#include <string.h>
#include <stdio.h>

int main(int argc, const char *argv[]) {
    char *c_str = NULL;
    size_t len;

    if (argc != 2) {
        printf("Usage: ./program string\n");
        exit(EXIT_FAILURE);
    }

    if ((len = strlen(argv[1])) >= 4) {
        c_str = (char *)malloc(len);

        if (!c_str) {
            perror("malloc");
        }

        strcpy(c_str, argv[1]);
        printf("%s\n", c_str);
    } else {
        c_str = "Some Literal String";
        printf("%s\n", c_str);
    }
    free(c_str);

    exit(EXIT_SUCCESS);
}