package com.xws.agentska.controller;

import com.xws.agentska.dto.UserDto;
import com.xws.agentska.dto.UserTokenState;
import com.xws.agentska.model.ContactInfo;
import com.xws.agentska.model.Role;
import com.xws.agentska.model.User;
import com.xws.agentska.security.TokenBasedAuthentication;
import com.xws.agentska.security.util.TokenUtils;
import com.xws.agentska.service.interfaces.UsersService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;

@RestController
public class UserController {
    @Autowired
    private UsersService usersService;
    @Autowired
    private AuthenticationManager authenticationManager;
    @Autowired
    private TokenUtils tokenUtils;
    @PostMapping(value = "/api/users", consumes = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<?> save(@RequestBody UserDto dto) {
        usersService.save(mapToUser(dto));
        return ResponseEntity.ok().build();
    }
    @PutMapping(value ="/api/users/session", consumes = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<?> login(@RequestBody UserDto dto) {
        Authentication authentication = authenticationManager
                .authenticate(new UsernamePasswordAuthenticationToken(dto.getUsername(), dto.getPassword()));
        SecurityContextHolder.getContext().setAuthentication(authentication);

        User user = (User) authentication.getPrincipal();
        String jwt = tokenUtils.generateToken(user);
        return ResponseEntity.ok(new UserTokenState(jwt, user.getRole().getName(),user.getId()));
    }
    private User mapToUser(UserDto dto){
        return new User(dto.getPassword(), dto.getName(), dto.getSurname(), new Role(dto.getRole()),dto.getUsername(), new ContactInfo(dto.getEmail(),dto.getPhone()));
    }
    public static long getLoggedinUserId() {
        return ((User) ((TokenBasedAuthentication) SecurityContextHolder.getContext().getAuthentication())
                .getPrincipal()).getId();
    }
}
