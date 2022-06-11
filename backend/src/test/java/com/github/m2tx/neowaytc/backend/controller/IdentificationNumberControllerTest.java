package com.github.m2tx.neowaytc.backend.controller;

import static org.hamcrest.CoreMatchers.any;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.delete;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.put;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import java.util.UUID;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.ValueSource;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.web.servlet.MockMvc;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.m2tx.neowaytc.backend.model.IdentificationNumber;

@SpringBootTest
@AutoConfigureMockMvc
@ActiveProfiles("test")
public class IdentificationNumberControllerTest {
    
	@Autowired
	private MockMvc mockMvc;

	@Autowired
	private ObjectMapper mapper;

    @ParameterizedTest
    @ValueSource(strings = {
        "39264468030",
        "10117286052",
		"02523798000181",
        "93566231000148",
        "89168173000127",
        "73953500000111"
    })
	public void addValidIdentificationNumber(String number) throws JsonProcessingException, Exception {
		this.mockMvc.perform(post("/")
				.contentType(MediaType.APPLICATION_JSON)
				.content(mapper.writeValueAsString(IdentificationNumber.builder().number(number).build())))
		.andDo(print())
		.andExpect(status().isCreated())
		.andExpect(jsonPath("$.number").value(number))
		.andExpect(jsonPath("$.id").isNotEmpty());
	}

	@ParameterizedTest
    @ValueSource(strings = {
        "046.847.189-81",
		"89927530000194",
        "200B",
        "1234"
    })
	public void failedToAddInvalidIdentificationNumber(String number) throws JsonProcessingException, Exception {
		this.mockMvc.perform(post("/")
				.contentType(MediaType.APPLICATION_JSON)
				.content(mapper.writeValueAsString(IdentificationNumber.builder().number(number).build())))
		.andDo(print())
		.andExpect(status().isBadRequest())
		.andExpect(jsonPath("$[0]").value("Invalid number: must be CPF or CNPJ"));
	}

}
