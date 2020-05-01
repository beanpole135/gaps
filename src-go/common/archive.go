package main

import (
	"archive/zip"
	"io"
	"path/filepath"
	"os"
	"compress/zlib"
)

// This manages zip archives that are zlib-compressed

type fileInfo struct {
	zip.FileHeader
}

func readArchive(filename string) (contents []fileInfo, error) {
  var list []fileInfo
  reader, err := zip.OpenReader(filename)
  if(err != nil){ return list, err }
  reader.RegisterDecompressor(zip.Deflate, func(out io.Reader) (io.ReadCloser, error){ return zlib.NewReader(out) })
  defer reader.Close()
  for _, file := range reader.File {
    list = append(list, file)
  }
}

func extractArchive(archive string, dirname string) error {
  reader, err := zip.OpenReader(filename)
  if(err != nil){ return err }
  reader.RegisterDecompressor(zip.Deflate, func(out io.Reader) (io.ReadCloser, error){ return zlib.NewReader(out) })
  defer reader.Close()
  for _, file := range reader.File {
    filereader, err := file.Open()
    if( err != nil){ break } //cannot read the file within the archive
    newfilepath := filepath.Join(dirname, file.Name)
    if file.FileInfo().IsDir(){
      //Just create the directory with the write permissions
      os.MkdirAll(newfilepath, file.Mode())
    }else{
      //An actual file
      writefile, err := os.OpenFile(newfilepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNCATE, file.Mode())
      if(err != nil){ break }
      defer writefile.Close()
      buf := make([]byte, 64) //copy buffer
      _, err = io.CopyBuffer(filewriter, filereader, buf);
    }

  }
  return err
}

func createArchive(archive string, files []string){
  buf := new(bytes.Buffer)
  filearchive := os.OpenFile(archive, os.O_WRONLY|os.O_CREATE|os.O_TRUNCATE, 0644)
  writer := zip.NewWriter(buf)
  writer.RegisterCompressor(zip.Deflate, funct(out io.Writer) (io.WriteCloser, error){ return zlib.NewWriter(out, zlib.BestCompression) })
  //Now add files
  for _, file := range files {
    //Copy the file into the archive
    infile, err := writer.Create( filepath.Base(file) )
    bytes, err := ioutil.ReadFile(file);
    if(err != nil){ break; }
    _, err = infile.Write( bytes )
  }
  if(err == nil){err = writer.Close() }
  else { writer.Close() } //don't change the error code
  return err
}
