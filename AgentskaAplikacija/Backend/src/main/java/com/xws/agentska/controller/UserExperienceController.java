package com.xws.agentska.controller;

import com.xws.agentska.dto.CommentDto;
import com.xws.agentska.dto.InterviewExperienceDto;
import com.xws.agentska.dto.SalaryExperienceDto;
import com.xws.agentska.dto.UserExperienceDto;
import com.xws.agentska.mapper.CustomModelMapper;
import com.xws.agentska.mapper.UserExperienceModelMapper;
import com.xws.agentska.model.*;
import com.xws.agentska.service.interfaces.UserExperienceService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import static com.xws.agentska.controller.UserController.getLoggedinUserId;

@RestController
public class UserExperienceController {

    @Autowired
    private UserExperienceService service;
    @Autowired
    private UserExperienceModelMapper userExperienceModelMapper;

    @GetMapping("/api/companies/{id}/comments")
    public ResponseEntity<?> findAllCommentsByCompany(@PathVariable int id) {
        return ResponseEntity.ok(service.findAllCommentsByCompanyId(id));
    }

    @GetMapping("/api/companies/{id}/salaries")
    public ResponseEntity<?> findAllSalariesByCompany(@PathVariable int id) {
        return ResponseEntity.ok(service.findAllSalariesByCompanyId(id));
    }

    @GetMapping("/api/companies/{id}/interviews")
    public ResponseEntity<?> findAllInterviewsByCompany(@PathVariable int id) {
        return ResponseEntity.ok(service.findAllInterviewsByCompanyId(id));
    }

    @PostMapping("/api/companies/{id}/comments")
    public ResponseEntity<?> createNewComment(@RequestBody CommentDto comment) {
        saveUserExperience(userExperienceModelMapper.convertToEntity(comment, Comment.class));
        return null;
    }

    @PostMapping("/api/companies/{id}/salaries")
    public ResponseEntity<?> createNewSalaryExperience(@RequestBody SalaryExperienceDto salaryExperience) {
        saveUserExperience(userExperienceModelMapper.convertToEntity(salaryExperience, SalaryExperience.class));
        return ResponseEntity.ok().build();
    }

    @PostMapping("/api/companies/{id}/interviews")
    public ResponseEntity<?> createNewInterviewExperience(@RequestBody InterviewExperienceDto interviewExperience) {
        saveUserExperience(userExperienceModelMapper.convertToEntity(interviewExperience, InterviewExperience.class));
        return ResponseEntity.ok().build();
    }

    private void saveUserExperience(UserExperience comment) {
        this.setLoggedinUserForCreator(comment);
        service.save(comment);
    }

    private void setLoggedinUserForCreator(UserExperience experience) {
        experience.setCreator(new User(getLoggedinUserId()));
    }

}
