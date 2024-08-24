import 'package:equatable/equatable.dart';

import 'user.dart';

class Product extends Equatable {
  final String id;
  final String name;
  final String description;
  final double price;
  final String imageUrl;
  final User seller;

  const Product(
      {required this.id,
      required this.name,
      required this.description,
      required this.price,
      required this.imageUrl,
      required this.seller})
      : super();

  @override
  List<Object?> get props => [id, name, description, price, imageUrl, seller];
}
