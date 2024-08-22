part of 'homepage_bloc.dart';

@immutable
sealed class HomepageEvent {}

class FetchProducts extends HomepageEvent {}

class AddButtonPress extends HomepageEvent {
  final String name;
  final String imageUrl;
  final double price;
  final String description;
  AddButtonPress(
      {required this.name,
      required this.imageUrl,
      required this.price,
      required this.description});
}

class DeleteButtonPress extends HomepageEvent {
  final String id;
  DeleteButtonPress({required this.id});
}

class UpdateButtonPress extends HomepageEvent {
  final String id;
  final String name;
  final String imageUrl;
  final double price;
  final String description;
  UpdateButtonPress(
      {required this.id,
      required this.name,
      required this.imageUrl,
      required this.price,
      required this.description});
}

// class SearchButtonPress extends HomepageEvent {
//   final String name;
//   SearchButtonPress({required this.name});
// }

