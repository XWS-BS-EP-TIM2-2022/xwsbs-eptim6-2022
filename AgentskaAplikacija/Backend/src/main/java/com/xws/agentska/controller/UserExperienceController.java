package com.xws.agentska.controller;

import com.xws.agentska.dto.CommentDto;
import com.xws.agentska.mapper.CommentModelMapper;
import com.xws.agentska.mapper.CustomModelMapper;
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
    private CustomModelMapper<Comment,CommentDto> commentModelMapper;
    @PostMapping("/api/comments")
    public ResponseEntity<?> createNewComment(@RequestBody CommentDto comment) {
        saveUserExperience(commentModelMapper.convertToEntity(comment));
        return null;
    }

    @GetMapping("/api/company/{id}/comments/")
    public ResponseEntity<?> findAllCommentsByCompany(@PathVariable int id) {
      return ResponseEntity.ok(service.findAllCommentsByCompanyId(id));
    }

    @PostMapping("/api/saleries")
    public ResponseEntity<?> createNewSalleryExperience(@RequestBody SalleryExperience salleryExperience) {
        saveUserExperience(salleryExperience);
        return ResponseEntity.ok().build();
    }

    @PostMapping("/api/interviews")
    public ResponseEntity<?> createNewInterviewExperience(@RequestBody InterviewExperience interviewExperience) {
        saveUserExperience(interviewExperience);
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
