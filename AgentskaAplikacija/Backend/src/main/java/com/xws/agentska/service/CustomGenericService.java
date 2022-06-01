package com.xws.agentska.service;

import com.xws.agentska.service.interfaces.Service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public class CustomGenericService<T,ID> implements Service<T,ID> {
    @Autowired
    protected JpaRepository<T, ID> repository;
    @Override
    public void save(T item) {
        repository.save(item);
    }

    @Override
    public List<T> findAll() {
        return repository.findAll();
    }

    @Override
    public T findById(ID id) {
        return repository.findById(id).get();
    }

    @Override
    public void update(T item) {
        repository.save(item);
    }

    @Override
    public void delete(ID id) {
        repository.deleteById(id);
    }
}
