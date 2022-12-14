module half_adder(  
    output S,C,  
    input A,B  
    );  
xor(S,A,B);  
and(C,A,B);  
endmodule  
 
module full_adder(  
    output S,Cout,  
    input A,B,Cin  
    );  
wire s1,c1,c2;  
half_adder HA1(s1,c1,A,B);  
half_adder HA2(S,c2,s1,Cin);  
or OG1(Cout,c1,c2);  
 
endmodule  
 
module ripple_adder_4bit(  
    output [3:0] Sum,  
    output Cout,  
    input [3:0] A,B,  
    input Cin  
    );  
wire c1,c2,c3;  
full_adder FA1(Sum[0],c1,A[0],B[0],Cin),  
FA2(Sum[1],c2,A[1],B[1],c1),  
FA3(Sum[2],c3,A[2],B[2],c2),  
FA4(Sum[3],Cout,A[3],B[3],c3);  
 
endmodule
