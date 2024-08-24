import 'package:equatable/equatable.dart';

class Seller extends Equatable {
  final String id;
  final String name;
  final String email;

  const Seller({
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
