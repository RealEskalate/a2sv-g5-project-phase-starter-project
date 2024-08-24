import 'package:equatable/equatable.dart';

class UserEntity extends Equatable {
  final String email;
  final String name;
  final String id;
  

  const UserEntity({
    required this.id,
    required this.email,
    required this.name,
  });

  @override
  List<Object?> get props => [email, name, id];
}
