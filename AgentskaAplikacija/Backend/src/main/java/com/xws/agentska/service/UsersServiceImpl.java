package com.xws.agentska.service;

import com.xws.agentska.model.Role;
import com.xws.agentska.model.User;
import com.xws.agentska.repository.RoleRepository;
import com.xws.agentska.repository.UsersRepository;
import com.xws.agentska.service.interfaces.UsersService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

import javax.transaction.Transactional;

@Service
public class UsersServiceImpl extends CustomGenericService<User,Long> implements UsersService {
    @Autowired
    private RoleRepository roleRepository;

    @Override
    public void save(User item) {
        Role role=roleRepository.findByName(item.getRole().getName());
        item.setRole(role);
        repository.save(item);
    }

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        return ((UsersRepository)repository).getByUsername(username);
    }
    @Transactional
    @Override
    public void updateUserRole(long userId, Role newRole) {
        User user = this.findById(userId);
        user.setRole(roleRepository.findByName(newRole.getName()));
        this.update(user);
    }
}
