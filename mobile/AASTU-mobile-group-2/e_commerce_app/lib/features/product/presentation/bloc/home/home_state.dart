part of 'home_bloc.dart';

//add sealed maybe

abstract class  HomeState  {
  

  
}

class HomeProductLoading extends HomeState {
  HomeProductLoading();
}

class HomeSuccessLoading extends HomeState {
  final List<ProductEntity> allProducts;
  HomeSuccessLoading({required this.allProducts});
}

class HomeFailureLoading extends HomeState {
 HomeFailureLoading();
}
