package main

import (
	"os"
	"os/exec"
	"fmt"
)

func exit_err(err error, details string){
  if(err != nil){
    fmt.Println("[ERROR] ", details);
    fmt.Println("  ", err)
    os.Exit(1)
  }
}

func doClean(subdir string){
  fmt.Println("Cleaning:", subdir)
  //First remove the compiled binary (if it exists)
  if _, err := os.Stat(subdir+"/"+subdir) ; !os.IsExist(err) {
    os.Remove(subdir+"/"+subdir)
  }
  //Now remove the link to the "common" directory (if it exists)
  if _, err := os.Stat(subdir+"/common") ; !os.IsExist(err) {
    os.Remove(subdir+"/common")
  }
}

func doPackage(subdir string) error {
  //Loop over all the OS/ARCH combinations and put the resulting binaries into specific dist dirs
  var err error
  err = nil


  return err
}

func doBuild(subdir string, OS string, ARCH string) error {
  fmt.Println("Starting build:", subdir)
  if( OS != "" && ARCH != ""){ fmt.Println("  OS:", OS, "ARCH:", ARCH) }
  //ensure the common dir is linked into the project
  if _, err := os.Stat(subdir+"/common") ; os.IsNotExist(err) {
    fmt.Println(" - Creating symlink to common module")
    os.Symlink("common", subdir+"/common");
  }
  cmd := exec.Command("go", "build", "-o", subdir)
    cmd.Dir = subdir
    if(OS != "") { cmd.Env = append(cmd.Env, "GOOS="+OS) }
    if(ARCH != "") { cmd.Env = append(cmd.Env, "GOARCH="+ARCH) }
  infoPipe, _ := cmd.CombinedOutput()
  err := cmd.Run()
  fmt.Printf("%s\n", infoPipe)
  return err
}

// This is the build routine for each of the projects
func main(){
  subcmd := "build"
  tools := []string{"gaps","gaps-pkg", "gaps-repo"}

  if(len(os.Args)==2){ subcmd = os.Args[1] }
  fmt.Println("Got subcmd:", subcmd);
  //Now run the appropriate type of operation
  var err error
  err = nil
  for _,tool := range(tools) {
    switch(subcmd){
      case "build":
        err = doBuild(tool,"","")

      case "clean":
        doClean(tool)

      case "package":
        err = doPackage(tool)

      default:
        fmt.Println("Unknown action: ", subcmd)
	fmt.Println("Available actions are:")
	fmt.Println(" - build:", "Compile the tools for the current system OS/ARCH")
	fmt.Println(" - clean:", "Cleanup all the build files")
	fmt.Println(" - package:", "Compile the tools for all os/architecture")
	fmt.Println(" - distclean:", "Cleanup all the package files")
        os.Exit(1)
    }
    if(err != nil){ os.Exit(1) }
  }
  os.Exit(0)
}
