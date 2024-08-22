import 'dart:async';

import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/domain/usecase/get_all_product_usecase.dart';
import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
part 'home_state.dart';
part 'home_event.dart';

class HomeBloc extends Bloc<HomeEvent, HomeState> {
  GetAllProductUsecase getAllProductUsecase;
  HomeBloc({required this.getAllProductUsecase}) : super(HomeProductLoading()) {
    on<HomeLoaded>(_onLoaded);
    on<HomeInitial>(_onInitial);
  }

  FutureOr<void> _onLoaded(HomeLoaded event, Emitter<HomeState> emit) async {
    emit(HomeProductLoading());

    final result = await getAllProductUsecase.execute();

    result.fold((failure) => emit(HomeFailureLoading()),
        (products) => emit(HomeSuccessLoading(allProducts: products)));
  }

  FutureOr<void> _onInitial(HomeInitial event, Emitter<HomeState> emit) {
    emit(HomeProductLoading());
  }
}
