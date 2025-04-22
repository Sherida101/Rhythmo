/*
 Rhythmo — Build better habits with rhythm
 https://github.com/Sherida101/Rhythmo

 See LICENSE for copyright details.
*/
package com.rhythmo.app.model;

import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Column;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import java.time.OffsetDateTime; // Handles datetime

@Entity
public class Habit {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY) // Specifies the generation strategy for the primary key
    private Long id;

    private String name;
    private String description;
    private String category;
    private String status;
    private String priority;
    private String frequency;
    private String reminder;
    private int streak;

    // @Column Annotation is used to specify the column details in the database
    @Column(name = "start_date", columnDefinition = "TIMESTAMP WITH TIME ZONE")
    private OffsetDateTime startDate;

    @Column(name = "end_date", columnDefinition = "TIMESTAMP WITH TIME ZONE")
    private OffsetDateTime endDate;

    @Column(name = "last_done", columnDefinition = "TIMESTAMP WITH TIME ZONE")
    private OffsetDateTime lastDone;

    @Column(name = "created_at", updatable = false, columnDefinition = "TIMESTAMP WITH TIME ZONE")
    private OffsetDateTime createdAt;

    @Column(name = "updated_at", columnDefinition = "TIMESTAMP WITH TIME ZONE")
    private OffsetDateTime updatedAt;

    // Constructors, getters, setters, etc.
    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getCategory() {
        return category;
    }

    public void setCategory(String category) {
        this.category = category;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public String getPriority() {
        return priority;
    }

    public void setPriority(String priority) {
        this.priority = priority;
    }

    public String getFrequency() {
        return frequency;
    }

    public void setFrequency(String frequency) {
        this.frequency = frequency;
    }

    public String getReminder() {
        return reminder;
    }

    public void setReminder(String reminder) {
        this.reminder = reminder;
    }

    public int getStreak() {
        return streak;
    }

    public void setStreak(int streak) {
        this.streak = streak;
    }
    public OffsetDateTime getStartDate() {
        return startDate;
    }
    public void setStartDate(OffsetDateTime startDate) {
        this.startDate = startDate;
    }
    public OffsetDateTime getEndDate() {
        return endDate;
    }
    public void setEndDate(OffsetDateTime endDate) {
        this.endDate = endDate;
    }
    public OffsetDateTime getLastDone() {
        return lastDone;
    }

    public void setLastDone(OffsetDateTime lastDone) {
        this.lastDone = lastDone;
    }

    public OffsetDateTime getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt() {
        this.createdAt = OffsetDateTime.now();
    }
    public OffsetDateTime getUpdatedAt() {
        return updatedAt;
    }

    public void setUpdatedAt() {
        this.updatedAt = OffsetDateTime.now();
    }

    // Constructor
    public Habit() {
        this.createdAt = OffsetDateTime.now();
        this.updatedAt = OffsetDateTime.now();
    }

    // Constructor with parameters
    public Habit(String name, String description, String category, String status, String priority, String frequency,
            String reminder) {
        this.name = name;
        this.description = description;
        this.category = category;
        this.status = status;
        this.priority = priority;
        this.frequency = frequency;
        this.reminder = reminder;
        this.streak = 0; // Default streak
        this.createdAt = OffsetDateTime.now();
        this.updatedAt = OffsetDateTime.now();
    }

    // toString method for easy debugging
    // This method is used to convert the object to a string representation
    @Override
    public String toString() {
        return "Habit{" +
                "id=" + id +
                ", name='" + name + '\'' +
                ", description='" + description + '\'' +
                ", category='" + category + '\'' +
                ", status='" + status + '\'' +
                ", priority='" + priority + '\'' +
                ", frequency='" + frequency + '\'' +
                ", reminder='" + reminder + '\'' +
                ", streak=" + streak +
                ", startDate=" + startDate +
                ", endDate=" + endDate +
                ", lastDone=" + lastDone +
                ", createdAt=" + createdAt +
                ", updatedAt=" + updatedAt +
                '}';
    }

    // Equals and hashCode methods can be added for better comparison and hashing
    @Override
    public boolean equals(Object o) {
        if (this == o)
            return true;
        if (!(o instanceof Habit))
            return false;

        Habit habit = (Habit) o;

        if (streak != habit.streak)
            return false;
        if (!id.equals(habit.id))
            return false;
        if (!name.equals(habit.name))
            return false;
        if (!description.equals(habit.description))
            return false;
        if (!category.equals(habit.category))
            return false;
        if (!status.equals(habit.status))
            return false;
        if (!priority.equals(habit.priority))
            return false;
        if (!frequency.equals(habit.frequency))
            return false;
        if (!reminder.equals(habit.reminder))
            return false;
        if (!startDate.equals(habit.startDate))
            return false;
        if (!endDate.equals(habit.endDate))
            return false;
        if (!lastDone.equals(habit.lastDone))
            return false;
        if (!createdAt.equals(habit.createdAt))
            return false;
        return updatedAt.equals(habit.updatedAt);
    }

    @Override
    public int hashCode() {
        int result = id.hashCode();
        result = 31 * result + name.hashCode();
        result = 31 * result + description.hashCode();
        result = 31 * result + category.hashCode();
        result = 31 * result + status.hashCode();
        result = 31 * result + priority.hashCode();
        result = 31 * result + frequency.hashCode();
        result = 31 * result + reminder.hashCode();
        result = 31 * result + streak;
        result = 31 * result + startDate.hashCode();
        result = 31 * result + endDate.hashCode();
        result = 31 * result + lastDone.hashCode();
        result = 31 * result + createdAt.hashCode();
        result = 31 * result + updatedAt.hashCode();
        return result;
    }
}
