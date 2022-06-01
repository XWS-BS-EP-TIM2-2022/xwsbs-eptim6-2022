package com.xws.agentska.repository;

import com.xws.agentska.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface UsersRepository extends JpaRepository<User, Long> {

    @Query("select u from User u where u.username = ?1")
    public User getByUsername(String username);
}