package com.kyriosdata.assinador;

import com.kyriosdata.assinador.domain.SignRequest;
import com.kyriosdata.assinador.domain.SignatureResponse;
import com.kyriosdata.assinador.domain.ValidateRequest;

public class Main {

    public static void main(String[] args) {
        if (args.length < 1) {
            System.err.println("Uso: assinador <sign|validate> [--content <valor>] [--token <valor>] [--signature <valor>]");
            System.exit(2);
        }

        String operation = args[0];

        if ("server".equals(operation)) {
            int port = 8080;
            int timeout = 0;
            for (int i = 1; i < args.length; i++) {
                if ("--port".equals(args[i]) && i + 1 < args.length) port = Integer.parseInt(args[++i]);
                if ("--timeout".equals(args[i]) && i + 1 < args.length) timeout = Integer.parseInt(args[++i]);
            }
            try {
                new AssinadorServer(port, timeout).start();
                Thread.currentThread().join();
            } catch (Exception e) {
                System.err.println("Erro ao iniciar servidor: " + e.getMessage());
                System.exit(1);
            }
            return;
        }

        String content = null;
        String token = null;
        String signature = null;

        for (int i = 1; i < args.length; i++) {
            switch (args[i]) {
                case "--content":
                    if (i + 1 < args.length) content = args[++i];
                    break;
                case "--token":
                    if (i + 1 < args.length) token = args[++i];
                    break;
                case "--signature":
                    if (i + 1 < args.length) signature = args[++i];
                    break;
                default:
                    System.err.println("Parâmetro desconhecido: " + args[i]);
                    System.exit(2);
            }
        }

        FakeSignatureService service = new FakeSignatureService();
        SignatureResponse response;

        switch (operation) {
            case "sign":
                response = service.sign(new SignRequest(content, token));
                break;
            case "validate":
                response = service.validate(new ValidateRequest(content, signature));
                break;
            default:
                System.err.println("Operação desconhecida: " + operation + ". Use 'sign', 'validate' ou 'server'.");
                System.exit(2);
                return;
        }

        System.out.println(response.toJson());
    }
}
