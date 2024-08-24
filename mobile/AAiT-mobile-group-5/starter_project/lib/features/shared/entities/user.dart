import 'package:equatable/equatable.dart';

class User extends Equatable {
  final String name;
  final String email;
  final String id;

  const User({
    required this.id,
    required this.name,
    required this.email,
  });

  @override
  List<Object?> get props => [
        id,
        name,
        email,
      ];
}
