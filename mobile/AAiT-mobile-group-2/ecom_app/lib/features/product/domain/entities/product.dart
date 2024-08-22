import 'package:equatable/equatable.dart';

class Product extends Equatable {
  final String id;
  final String name;
  final String description;
  final String imageUrl;
  final double price;

  const Product(
      {required this.id,
      required this.name,
      required this.description,
      required this.imageUrl,
      required this.price});

  @override
  List<Object?> get props => [id, name, description, imageUrl, price];
}
