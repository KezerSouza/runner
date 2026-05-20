package jdk

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestIs21Plus_NonExistent(t *testing.T) {
	if Is21Plus("/nao/existe/java") {
		t.Error("Is21Plus deve retornar false para binário inexistente")
	}
}

func TestIs21Plus_NotJava(t *testing.T) {
	tmp := t.TempDir()
	fake := filepath.Join(tmp, Exe())
	if err := os.WriteFile(fake, []byte("#!/bin/sh\necho 'not java'"), 0755); err != nil {
		t.Fatal(err)
	}
	if Is21Plus(fake) {
		t.Error("Is21Plus deve retornar false para script sem saída de versão Java")
	}
}

func TestIs21Plus_SystemJava(t *testing.T) {
	java, err := exec.LookPath(Exe())
	if err != nil {
		t.Skipf("java não encontrado no PATH: %v", err)
	}
	if !Is21Plus(java) {
		t.Errorf("Is21Plus(%q) = false; ambiente requer Java 21+", java)
	}
}

func TestFindIn_EmptyDir(t *testing.T) {
	tmp := t.TempDir()
	_, ok := findIn(tmp)
	if ok {
		t.Error("findIn em diretório vazio deve retornar false")
	}
}

func TestFindIn_DirWithoutJava(t *testing.T) {
	tmp := t.TempDir()
	if err := os.MkdirAll(filepath.Join(tmp, "jdk-21", "bin"), 0755); err != nil {
		t.Fatal(err)
	}
	_, ok := findIn(tmp)
	if ok {
		t.Error("findIn sem binário java deve retornar false")
	}
}

func TestExe(t *testing.T) {
	if Exe() == "" {
		t.Error("Exe() não deve retornar string vazia")
	}
}
