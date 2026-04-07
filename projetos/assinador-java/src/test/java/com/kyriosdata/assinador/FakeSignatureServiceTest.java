package com.kyriosdata.assinador;

import com.kyriosdata.assinador.domain.SignRequest;
import com.kyriosdata.assinador.domain.SignatureResponse;
import com.kyriosdata.assinador.domain.ValidateRequest;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class FakeSignatureServiceTest {

    private FakeSignatureService service;

    @BeforeEach
    void setUp() {
        service = new FakeSignatureService();
    }

    @Test
    void deveRetornarAssinaturaSimuladaParaRequisicaoValida() {
        SignRequest request = new SignRequest("conteudo do documento", null);
        SignatureResponse response = service.sign(request);

        assertTrue(response.isValid());
        assertEquals(FakeSignatureService.FAKE_SIGNATURE, response.getSignature());
    }

    @Test
    void deveRejeitarRequisicaoDeAssinaturaComConteudoNulo() {
        SignRequest request = new SignRequest(null, null);
        SignatureResponse response = service.sign(request);

        assertFalse(response.isValid());
        assertNotNull(response.getMessage());
    }

    @Test
    void deveRejeitarRequisicaoDeAssinaturaComConteudoVazio() {
        SignRequest request = new SignRequest("  ", null);
        SignatureResponse response = service.sign(request);

        assertFalse(response.isValid());
    }

    @Test
    void deveValidarAssinaturaSimuladaCorretamente() {
        ValidateRequest request = new ValidateRequest("conteudo", FakeSignatureService.FAKE_SIGNATURE);
        SignatureResponse response = service.validate(request);

        assertTrue(response.isValid());
    }

    @Test
    void deveRejeiarAssinaturaInvalida() {
        ValidateRequest request = new ValidateRequest("conteudo", "assinatura-errada");
        SignatureResponse response = service.validate(request);

        assertFalse(response.isValid());
    }

    @Test
    void deveRejeitarValidacaoComConteudoNulo() {
        ValidateRequest request = new ValidateRequest(null, FakeSignatureService.FAKE_SIGNATURE);
        SignatureResponse response = service.validate(request);

        assertFalse(response.isValid());
        assertNotNull(response.getMessage());
    }

    @Test
    void deveRejeitarValidacaoComAssinaturaNula() {
        ValidateRequest request = new ValidateRequest("conteudo", null);
        SignatureResponse response = service.validate(request);

        assertFalse(response.isValid());
        assertNotNull(response.getMessage());
    }
}
