/*
 Rhythmo — Build better habits with rhythm
 https://github.com/Sherida101/Rhythmo

 See LICENSE for copyright details.
*/
package com.rhythmo.app.repository;

import java.util.List;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import com.rhythmo.app.model.Habit;

@Repository
public interface HabitRepository extends JpaRepository<Habit, Long> {
    // Create

    // Delete

    // Get
    // No need to define for findAll() because it is built-in
    List<Habit> findByUserId(Long userId);
    List<Habit> findByCategory(String category);
    List<Habit> findByFrequency(String frequency);
    List<Habit> findByStatus(String status);
    List<Habit> findByTag(String tag);
    List<Habit> findByPriority(String priority);

    // Update
}
