package com.github.m2tx.neowaytc.backend.validator;

import static java.lang.annotation.ElementType.FIELD;
import static java.lang.annotation.ElementType.PARAMETER;
import static java.lang.annotation.RetentionPolicy.RUNTIME;

import java.lang.annotation.Documented;
import java.lang.annotation.Retention;
import java.lang.annotation.Target;

import javax.validation.Constraint;
import javax.validation.Payload;

@Target( { FIELD, PARAMETER })
@Retention(RUNTIME)
@Documented
@Constraint(validatedBy = CpfCnpjValidator.class)
public @interface CpfCnpjValidation {    
    public String message() default "Invalid number: must be CPF or CNPJ";
    public Class<?>[] groups() default {};
    public Class<? extends Payload>[] payload() default {};
}