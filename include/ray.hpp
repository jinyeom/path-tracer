#ifndef RAY_HPP
#define RAY_HPP

#include <glm>

class Ray
{
private:
    glm::dvec3 p_; // origin point of the ray      
    glm::dvec3 d_; // direction of the ray

public:
    // Create a new ray given its origin point and its direction.
    Ray(glm::dvec3 p, glm::dvec3 d);

    glm::dvec3 getOrigin() const { return p_; }
    glm::dvec3 getDirection() const { return d_; }
};

#endif
