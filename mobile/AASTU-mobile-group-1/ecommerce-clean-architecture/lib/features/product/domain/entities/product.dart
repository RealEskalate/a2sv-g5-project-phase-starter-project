import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../auth/domain/entities/user.dart';

class Productentity extends Equatable {
  const Productentity({
    required this.id,
    required this.image,
    required this.name,
    required this.description,
    required this.price,
    required this.seller,
  });
  final String id;
  final String image;
  final String name;
  final String description;
  final double price;
  final UserEntity seller;

  @override
  List<Object?> get props => [
        id,
        image,
        name,
        description,
        price,
      ];
}
