package com.github.m2tx.neowaytc.backend.exceptions;

public class IdentificationNumberAlreadyExistsException extends Exception{
    
    public IdentificationNumberAlreadyExistsException(){
        super("Identification Number already exists!");
    }

    public IdentificationNumberAlreadyExistsException(String message){
        super(message);
    }

}
