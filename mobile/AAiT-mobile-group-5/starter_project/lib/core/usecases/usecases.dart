import 'package:dartz/dartz.dart'; // Importing the dartz package for functional programming and Either type.
import 'package:equatable/equatable.dart'; // Importing equatable for value equality checks.

import '../errors/failure.dart'; // Importing custom failure handling class.

/// Abstract class `UseCase` which represents a single unit of work in the app.
/// `Type` represents the type of data it returns (e.g., a list of products).
/// `Params` represents the parameters it needs to execute (e.g., an ID).
abstract class UseCase<Type, Params> {
  /// Abstract method `call` which must be implemented by subclasses.
  /// This method takes `Params` and returns a `Future` that either
  /// succeeds with `Type` or fails with `Failure`.
  Future<Either<Failure, Type>> call(Params params);
}

/// Class `NoParams` is used when a `UseCase` doesn't require any parameters.
/// It extends `Equatable` to enable value comparison, making it easier to manage state.
class NoParams extends Equatable {
  @override
  List<Object?> get props =>
      []; // Returns an empty list for equality comparison.
}
