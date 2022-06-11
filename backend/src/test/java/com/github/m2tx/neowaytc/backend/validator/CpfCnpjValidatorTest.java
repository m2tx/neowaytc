package com.github.m2tx.neowaytc.backend.validator;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.ValueSource;

public class CpfCnpjValidatorTest {
    
    @ParameterizedTest
    @ValueSource(strings = {
        "39264468030",
        "10117286052"
    })
    public void isValidCPF(String cpf) {
        assertTrue(CpfCnpjValidator.isValidCPF(cpf));
    }

    @ParameterizedTest
    @ValueSource(strings = {
        "02523798000181",
        "93566231000148",
        "89168173000127",
        "73953500000111"
    })
    public void isValidCNPJ(String cnpj) {
        assertTrue(CpfCnpjValidator.isValidCNPJ(cnpj));
    }

    @ParameterizedTest
    @ValueSource(strings = {
        "00000000000",
        "11111111111",
        "22222222222",
        "33333333333",
        "44444444444",
        "55555555555",
        "66666666666",
        "77777777777",
        "88888888888",
        "99999999999",
        "10117286051"
    })
    public void isInvalidCPF(String cpf) {
        assertFalse(CpfCnpjValidator.isValidCPF(cpf));
    }

    @ParameterizedTest
    @ValueSource(strings = {
        "00000000000000",
        "11111111111111",
        "22222222222222",
        "33333333333333",
        "44444444444444",
        "55555555555555",
        "66666666666666",
        "77777777777777",
        "88888888888888",
        "99999999999999",
        "89927530000194"
    })
    public void isInvalidCNPJ(String cnpj) {
        assertFalse(CpfCnpjValidator.isValidCNPJ(cnpj));
    }

    @ParameterizedTest
    @ValueSource(strings = {
        "767.152.770-01",
        "392.644.680-30",
        "101.172.860-52",
        "93.566.231/0001-48",
        "73.953.500/0001-11"
    })
    public void isValidFormatedValues(String value) {
        assertTrue(CpfCnpjValidator.isValid(value));
    }

    @ParameterizedTest
    @ValueSource(strings = {
        "046.847.189-81",
        "200B",
        "1234"
    })
    public void isInvalidOtherValues(String value) {
        assertFalse(CpfCnpjValidator.isValid(value));
    }

    @Test
    public void isInvalidNullValue() {
        assertFalse(CpfCnpjValidator.isValid(null));
    }
}
