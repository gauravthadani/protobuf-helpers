# protobuf-helpers

A collection of helper functions for working with Protocol Buffer objects in Go.

## Overview

This repository contains utility functions to simplify common operations when dealing with Protocol Buffers in Golang projects. These helpers aim to streamline tasks and improve efficiency when working with protobuf objects.

## Features

### Equality Check Ignoring Deprecated Fields

The primary utility function in this package allows for comparing two protobuf objects while ignoring any deprecated fields. This is particularly useful when you need to check for equality between objects that may have evolved over time, with some fields becoming deprecated.

## Installation

To use these helpers in your project, you can install the package using:
