package com.kyriosdata.assinador.domain;

public class SignatureResponse {
    private String signature;
    private boolean valid;
    private String message;

    public SignatureResponse(String signature, boolean valid, String message) {
        this.signature = signature;
        this.valid = valid;
        this.message = message;
    }

    public String getSignature() {
        return signature;
    }

    public boolean isValid() {
        return valid;
    }

    public String getMessage() {
        return message;
    }
}
