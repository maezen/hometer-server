/*
 * Hometer
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

import (
	"time"
)

type Pressure struct {
	Date time.Time `json:"date,omitempty"`

	Value float64 `json:"value,omitempty"`
}
