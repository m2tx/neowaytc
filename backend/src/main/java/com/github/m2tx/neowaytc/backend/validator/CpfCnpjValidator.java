package com.github.m2tx.neowaytc.backend.validator;

import static java.lang.Integer.parseInt;
import static org.apache.commons.lang3.StringUtils.reverse;

import javax.validation.ConstraintValidator;
import javax.validation.ConstraintValidatorContext;

public class CpfCnpjValidator implements ConstraintValidator<CpfCnpjValidation, String>{
    
    public static boolean isValid(String cpfCnpj) {
        if(cpfCnpj==null) return false;
        cpfCnpj = removeNotNumberCharacters(cpfCnpj);
        if (cpfCnpj.length() == 11) {
            return isValidCPF(cpfCnpj);
        } else if (cpfCnpj.length() == 14) {
            return isValidCNPJ(cpfCnpj);
        }
        return false;
    }

    protected static String removeNotNumberCharacters(String cpfCnpj) {
        return cpfCnpj.replaceAll("[^0-9]", "");
    }

    protected static boolean isValidCPF(String cpf) {
        if (cpf.equals("00000000000") || 
            cpf.equals("11111111111") || 
            cpf.equals("22222222222") || 
            cpf.equals("33333333333") || 
            cpf.equals("44444444444") || 
            cpf.equals("55555555555") || 
            cpf.equals("66666666666") || 
            cpf.equals("77777777777") || 
            cpf.equals("88888888888") || 
            cpf.equals("99999999999")) {
            return false;
        }
        int d1 = 0;
        for (int i = 0; i < 9; i++) {
            d1 += parseInt(cpf.substring(i, i + 1)) * (10 - i);
        }
        if (d1 == 0) {
            return false;
        }
        d1 = 11 - (d1 % 11);
        if (d1 > 9) {
            d1 = 0;
        }
        if (parseInt(cpf.substring(9, 10)) != d1) {
            return false;
        }
        d1 *= 2;
        for (int i = 0; i < 9; i++) {
            d1 += parseInt(cpf.substring(i, i + 1)) * (11 - i);
        }
        d1 = 11 - (d1 % 11);
        if (d1 > 9) {
            d1 = 0;
        }
        if (parseInt(cpf.substring(10, 11)) != d1) {
            return false;
        }
        return true;
    }

    protected static boolean isValidCNPJ(String cnpj) {
        if (cnpj.equals("00000000000000") || 
            cnpj.equals("11111111111111") || 
            cnpj.equals("22222222222222") || 
            cnpj.equals("33333333333333") || 
            cnpj.equals("44444444444444") || 
            cnpj.equals("55555555555555") || 
            cnpj.equals("66666666666666") || 
            cnpj.equals("77777777777777") || 
            cnpj.equals("88888888888888") || 
            cnpj.equals("99999999999999")) {
            return false;
        }      
        String reverseCnpj = reverse(cnpj);          
        int d1 = 0;
        int p = 2;
        for (int i = 2; i < 14; i++) {
            d1 += parseInt(reverseCnpj.substring(i, i + 1)) * p;            
            if(++p>9) p = 2;
        }
        if (d1 == 0) {
            return false;
        }
        d1 = 11 - (d1 % 11);
        if (d1 > 9) {
            d1 = 0;
        }
        if (parseInt(cnpj.substring(12, 13)) != d1) {
            return false;
        }
        d1 *= 2;
        p = 3;
        for (int i = 2; i < 14; i++) {
            d1 += parseInt(reverseCnpj.substring(i, i + 1)) * p;
            if(++p>9) p = 2;
        }
        d1 = 11 - (d1 % 11);
        if (d1 > 9) {
            d1 = 0;
        }
        if (parseInt(cnpj.substring(13, 14)) != d1) {
            return false;
        }
        return true;
    }

    @Override
    public boolean isValid(String value, ConstraintValidatorContext context) {
        return isValid(value);
    }

}
