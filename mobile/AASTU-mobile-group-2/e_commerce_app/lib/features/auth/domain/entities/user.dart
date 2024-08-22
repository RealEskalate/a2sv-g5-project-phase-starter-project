import 'package:equatable/equatable.dart';

class User extends Equatable {
  final String id;
  final String email;
  final String name;

  User({
    required this.id,
    required this.email,
    required this.name,


  });

  @override
  List<Object?> get props => [id, email, name];
}