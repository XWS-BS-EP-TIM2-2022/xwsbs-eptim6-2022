package com.xws.agentska.service.interfaces;

import java.util.List;

public interface Service<T,ID> {
	void save(T item);
	List<T> findAll();
	T findById(ID id);
	void update(T item);
	void delete(ID id);
}
