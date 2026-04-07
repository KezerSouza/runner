package com.kyriosdata.assinador;

import com.kyriosdata.assinador.domain.SignRequest;
import com.kyriosdata.assinador.domain.SignatureResponse;
import com.kyriosdata.assinador.domain.ValidateRequest;

public class FakeSignatureService implements SignatureService {

    static final String FAKE_SIGNATURE = "MOCKED_SIGNATURE_BASE64_==";

    @Override
    public SignatureResponse sign(SignRequest request) {
        if (request.getContent() == null || request.getContent().isBlank()) {
            return new SignatureResponse(null, false, "Parâmetro 'content' é obrigatório e não pode ser vazio.");
        }
        return new SignatureResponse(FAKE_SIGNATURE, true, "Assinatura simulada gerada com sucesso.");
    }

    @Override
    public SignatureResponse validate(ValidateRequest request) {
        if (request.getContent() == null || request.getContent().isBlank()) {
            return new SignatureResponse(null, false, "Parâmetro 'content' é obrigatório e não pode ser vazio.");
        }
        if (request.getSignature() == null || request.getSignature().isBlank()) {
            return new SignatureResponse(null, false, "Parâmetro 'signature' é obrigatório e não pode ser vazio.");
        }
        boolean valid = FAKE_SIGNATURE.equals(request.getSignature());
        String message = valid ? "Assinatura válida." : "Assinatura inválida.";
        return new SignatureResponse(null, valid, message);
    }
}
