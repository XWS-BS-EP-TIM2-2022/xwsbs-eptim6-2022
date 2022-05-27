package com.xws.agentska.mapper;

import java.util.List;

public interface CustomModelMapper<E,T> {
	public T convertToDto(E entity);
	public E convertToEntity(T dto);
	public List<T> convertToDtos(List<E> entities);
	public List<E> convertToEntities(List<T> dtos);
}
