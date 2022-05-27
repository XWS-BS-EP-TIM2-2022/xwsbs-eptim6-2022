package com.xws.agentska.service.interfaces;

import com.xws.agentska.model.Role;
import com.xws.agentska.model.User;
import com.xws.agentska.service.interfaces.Service;
import org.springframework.security.core.userdetails.UserDetailsService;

public interface UsersService extends UserDetailsService, Service<User,Long> {
    public void updateUserRole(long userId, Role newRole);
}