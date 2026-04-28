package simulador

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Arquivo que armazena o PID do processo
func stateFile() string {
	home, _ := os.UserHomeDir()

	dir := filepath.Join(home, ".hubsaude")
	os.MkdirAll(dir, 0755)

	return filepath.Join(dir, "simulador.pid")
}

// Caminho do simulador.jar
func jarPath() string {
	if p := os.Getenv("SIMULADOR_JAR"); p != "" {
		return p
	}

	return "simulador.jar"
}

// Caminho do Java instalado
func javaPath() string {
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		if runtime.GOOS == "windows" {
			return filepath.Join(jh, "bin", "java.exe")
		}

		return filepath.Join(jh, "bin", "java")
	}

	return "java"
}

// Verifica se a porta está livre
func portFree(port int) bool {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}

	ln.Close()
	return true
}

// Inicia o simulador
func Start(port int) error {
	if !portFree(port) {
		return fmt.Errorf("porta %d ocupada", port)
	}

	cmd := exec.Command(
		javaPath(),
		"-jar",
		jarPath(),
		"--server.port="+strconv.Itoa(port),
	)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("erro ao iniciar simulador: %w", err)
	}

	pid := strconv.Itoa(cmd.Process.Pid)

	return os.WriteFile(stateFile(), []byte(pid), 0644)
}

// Encerra o simulador
func Stop() error {
	data, err := os.ReadFile(stateFile())
	if err != nil {
		return fmt.Errorf("simulador não está em execução")
	}

	pid, _ := strconv.Atoi(strings.TrimSpace(string(data)))

	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	_ = proc.Kill()
	_ = os.Remove(stateFile())

	return nil
}

// Retorna status atual
func Status() (bool, int) {
	data, err := os.ReadFile(stateFile())
	if err != nil {
		return false, 0
	}

	pid, _ := strconv.Atoi(strings.TrimSpace(string(data)))

	return true, pid
}
